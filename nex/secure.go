package nex

import (
	"fmt"
	"os"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/tri-force-heroes/globals"
)

func StartSecureServer() {
	globals.SecureServer = nex.NewServer()
	globals.SecureServer.SetPRUDPVersion(1)
	globals.SecureServer.SetPRUDPProtocolMinorVersion(3)
	globals.SecureServer.SetDefaultNEXVersion(nex.NewNEXVersion(3,9,0))
	globals.SecureServer.SetKerberosPassword(globals.KerberosPassword)
	globals.SecureServer.SetAccessKey("c1621b84")

	globals.Timeline = make(map[uint32][]uint8)

	globals.SecureServer.On("Data", func(packet *nex.PacketV1) {
		request := packet.RMCRequest()

		fmt.Println("==Tri-Force Heroes - Secure==")
		fmt.Printf("Protocol ID: %#v\n", request.ProtocolID())
		fmt.Printf("Method ID: %#v\n", request.MethodID())
		fmt.Println("===============")
	})

	registerCommonSecureServerProtocols()
	registerSecureServerNEXProtocols()

	globals.SecureServer.Listen(fmt.Sprintf(":%s", os.Getenv("PN_TFH_SECURE_SERVER_PORT")))
}
