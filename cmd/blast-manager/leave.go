// Copyright (c) 2019 Minoru Osuka
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 		http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"os"

	"github.com/mosuka/blast/manager"
	"github.com/mosuka/blast/protobuf/raft"
	"github.com/urfave/cli"
)

func execLeave(c *cli.Context) error {
	peerAddr := c.String("peer-addr")
	nodeId := c.String("node-id")

	node := &raft.Node{
		Id: nodeId,
	}

	peerClient, err := manager.NewGRPCClient(peerAddr)
	if err != nil {
		return err
	}
	defer func() {
		err := peerClient.Close()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}()

	err = peerClient.Leave(node)
	if err != nil {
		return err
	}

	return nil
}
