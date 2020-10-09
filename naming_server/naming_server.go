package naming_server

import (
	"fmt"
	"google.golang.org/grpc"
	"math/rand"
	"net"
	"os"
	utils "project-dfs"
	"project-dfs/pb"
	"strings"
	"sync"
)

type StorageInfo struct {
	Alias string
}

type NodeType int32

const (
	DIR  NodeType = 1
	FILE NodeType = 2
)

type Node struct {
	Name     string // name of a file or directory
	Type     NodeType
	Children []*Node
	Storages []*StorageInfo
}

func (n *Node) GetChild(name string) *Node {
	for _, child := range n.Children {
		if child.Name == name {
			return child
		}
	}

	return nil
}

func (n *Node) RemoveChild(name string) {
	index := -1
	for i, child := range n.Children {
		if child.Name == name {
			index = i
		}
	}
	if index == -1 {
		println("Couldn't remove child"+name+"from node", n.Name, "as it doesn't exist")
		return
	}
	n.Children = append(n.Children[:index], n.Children[index+1:]...)
}

func (n *Node) AddChild(node *Node) {
	n.Children = append(n.Children, node)
}

func NewNode(name string, t NodeType) *Node {
	return &Node{
		Name:     name,
		Type:     t,
		Children: make([]*Node, 0),
		Storages: make([]*StorageInfo, 0),
	}
}

func (server *NamingServer) FindNode(path string) (*Node, bool) {
	node := server.RootIndexNode

	segments := strings.Split(path, "/")[1:]
	fmt.Println("Segments:", strings.Join(segments, ", "))

	for _, s := range segments {
		updated := false
		for _, child := range node.Children {
			if child.Name == s {
				node = child
				updated = true
				break
			}
		}
		if !updated {
			println("warning: no node found in FindNode for path " + path)
			return nil, false
		}
	}

	return node, true
}

func (server *NamingServer) CreateNodeIfNotExists(path string, lastNodeIsFile bool) *Node {
	segments := strings.Split(path, "/")[1:]
	node := server.RootIndexNode
	for _, s := range segments {
		exists := false
		for _, child := range node.Children {
			if child.Name == s {
				node = child
				exists = true
				break
			}
		}
		if !exists {
			t := DIR
			if lastNodeIsFile && s == segments[len(segments)-1] {
				t = FILE
			}

			n := &Node{
				Name:     s,
				Type:     t,
				Children: []*Node{},
				Storages: []*StorageInfo{},
			}
			node.Children = append(node.Children, n)
			node = n
		}
	}

	return node
}

type NamingServer struct {
	storageAddressesMutex sync.Mutex
	StorageAddresses      map[string]string // key:value = serverAlias:serverAddress
	LocalAddress          string
	RootIndexNode         *Node
	StorageServers        map[string]pb.StorageClient
}

func (server *NamingServer) SetAddressMap(newKey string, newValue string) {
	server.storageAddressesMutex.Lock()
	defer server.storageAddressesMutex.Unlock()
	server.StorageAddresses[newKey] = newValue
}

// Returns 2 random storage servers. That's it.
func (server *NamingServer) Get2RandomStorageServers() []*pb.DiscoveredStorage {
	servers := server.StorageAddresses
	keys := utils.Keys(servers)
	var result []*pb.DiscoveredStorage

	var aliases []string
	for {
		index := rand.Intn(len(servers))
		alias := keys[index]
		if utils.Contains(aliases, alias) {
			continue
		}
		aliases = append(aliases, alias)
		if len(aliases) == 2 {
			break
		}
	}

	for _, alias := range aliases {
		address := servers[alias]

		result = append(result, &pb.DiscoveredStorage{
			Alias:   alias,
			Address: address,
		})
	}

	return result
}

func (server *NamingServer) GetStorageServer(address string) pb.StorageClient {
	ss, ok := server.StorageServers[address]
	if !ok {
		conn, err := grpc.Dial(address, grpc.WithInsecure())
		if err != nil {
			println("GetStorageServer: error dialing storage server:", err.Error())
			return nil
		}
		ss = pb.NewStorageClient(conn)
		server.StorageServers[address] = ss
	}
	return ss
}

func initNamingServer() *NamingServer {
	// Obtain address from environment
	address := os.Getenv("ADDRESS")
	if address == "" {
		address = "0.0.0.0:5678"
		fmt.Println("ADDRESS variable not specified; falling back to", address)
	}

	rootNode := &Node{
		Name:     "",
		Children: make([]*Node, 0),
	}

	return &NamingServer{
		storageAddressesMutex: sync.Mutex{},
		StorageAddresses:      make(map[string]string),
		LocalAddress:          address,
		RootIndexNode:         rootNode,
	}
}

func Run() {
	server := initNamingServer()

	println("Initialized metadata: ")
	fmt.Printf("%+v\n", server)

	listener, err := net.Listen("tcp", server.LocalAddress)
	if err != nil {
		println("Error listening:", err.Error())
		os.Exit(0)
	}
	println("Listening on " + server.LocalAddress)

	namingController := NewNamingServiceController(server)
	grpcServer := grpc.NewServer()
	pb.RegisterNamingServer(grpcServer, namingController)
	err = grpcServer.Serve(listener)
	if err != nil {
		println("Error serving:", err.Error())
		os.Exit(0)
	}
}
