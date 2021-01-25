package main

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"
)


func tabPrinter() func(...interface{}) {
	w := tabwriter.NewWriter(os.Stdout, 33, 0, 0, ' ', 0)
	fmt.Fprintln(w, strings.Join([]string{"State", "Recv-Q", "Send-Q", "Local", "Foreign", "Timer", "TimerExp", "TimerRet"}, "\t"))
	printerFn := func(fields ...interface{}) { fmt.Fprintf(w, "%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\t\n", fields...); w.Flush() }
	return printerFn
}

func writeStateChanges(currentSocketDiag *socketDiagMsg, writerFn func(...interface{})) {
	var state, recvQ, sendQ, localAddress, remoteAddress, timer, timerExpiry, retransmissions string
	writeFieldsFn := func() {
		writerFn(state, recvQ, sendQ, localAddress, remoteAddress, timer, timerExpiry, retransmissions)
	}

	state = currentSocketDiag.state
	recvQ = fmt.Sprintf("%v", currentSocketDiag.recvQ)
	sendQ = fmt.Sprintf("%v", currentSocketDiag.sendQ)
	localAddress = currentSocketDiag.localAddress
	remoteAddress = currentSocketDiag.remoteAddress
	timer = currentSocketDiag.timer
	retransmissions = fmt.Sprintf("%v", currentSocketDiag.retransmissions)
	timerExpiry = fmt.Sprintf("%v", currentSocketDiag.timerExpiry)
	writeFieldsFn()
}
