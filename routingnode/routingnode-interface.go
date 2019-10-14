// jeffCoin routingnode-interface.go

package rountingnode

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"strconv"
	"sync"
	"time"

	errors "github.com/pkg/errors"

	log "github.com/sirupsen/logrus"
)

