package node

import (
	"context"
	"fmt"
	"sync"

	log "github.com/sirupsen/logrus"
	panel "github.com/wyx2685/v2node/api/v2board"
	"github.com/wyx2685/v2node/conf"
	"github.com/wyx2685/v2node/core"
)

type Node struct {
	controllers []*Controller
	NodeInfos   []*panel.NodeInfo
}

type nodeInitResult struct {
	index      int
	controller *Controller
	nodeInfo   *panel.NodeInfo
	err        error
}

func New(nodes []conf.NodeConfig) (*Node, error) {
	n := &Node{
		controllers: make([]*Controller, len(nodes)),
		NodeInfos:   make([]*panel.NodeInfo, len(nodes)),
	}

	if len(nodes) == 0 {
		return n, nil
	}

	resultCh := make(chan nodeInitResult, len(nodes))
	var wg sync.WaitGroup

	for i, nodeConfig := range nodes {
		wg.Add(1)
		go func(index int, config conf.NodeConfig) {
			defer wg.Done()
			p, err := panel.New(&config)
			if err != nil {
				resultCh <- nodeInitResult{index: index, err: err}
				return
			}
			info, err := p.GetNodeInfo(context.Background())
			if err != nil {
				resultCh <- nodeInitResult{index: index, err: err}
				return
			}
			controller := NewController(p, &config, info)
			resultCh <- nodeInitResult{index: index, controller: controller, nodeInfo: info}
		}(i, nodeConfig)
	}

	go func() {
		wg.Wait()
		close(resultCh)
	}()

	for result := range resultCh {
		if result.err != nil {
			return nil, result.err
		}
		n.controllers[result.index] = result.controller
		n.NodeInfos[result.index] = result.nodeInfo
	}

	return n, nil
}

func (n *Node) Start(nodes []conf.NodeConfig, core *core.V2Core) error {
	for i, node := range nodes {
		err := n.controllers[i].Start(core)
		if err != nil {
			return fmt.Errorf("start node controller [%s-%d] error: %s",
				node.APIHost,
				node.NodeID,
				err)
		}
	}
	return nil
}

func (n *Node) Close() error {
	var err error
	for _, c := range n.controllers {
		if err = c.Close(); err != nil {
			log.Errorf("close controller failed: %v", err)
			return err
		}
	}
	n.controllers = nil
	return nil
}
