package nex

import (
	"fmt"
	"os"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/tri-force-heroes/globals"
)

var serverBuildString string

func StartAuthenticationServer() {
	globals.AuthenticationServer = nex.NewServer()
	globals.AuthenticationServer.SetPRUDPVersion(1)
	globals.AuthenticationServer.SetPRUDPProtocolMinorVersion(2)
	globals.AuthenticationServer.SetDefaultNEXVersion(nex.NewNEXVersion(3,9,0))
	globals.AuthenticationServer.SetKerberosPassword(globals.KerberosPassword)
	globals.AuthenticationServer.SetAccessKey("c1621b84")

	globals.AuthenticationServer.On("Data", func(packet *nex.PacketV1) {
		request := packet.RMCRequest()

		fmt.Println("==Tri-Force Heroes - Auth==")
		fmt.Printf("Protocol ID: %#v\n", request.ProtocolID())
		fmt.Printf("Method ID: %#v\n", request.MethodID())
		fmt.Println("===============")
	})

	registerCommonAuthenticationServerProtocols()

	globals.AuthenticationServer.Listen(fmt.Sprintf(":%s", os.Getenv("PN_TFH_AUTHENTICATION_SERVER_PORT")))
}
