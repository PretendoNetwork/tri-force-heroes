package nex

import (
	utility "github.com/PretendoNetwork/nex-protocols-go/utility"
	"github.com/PretendoNetwork/tri-force-heroes/globals"
)

func registerSecureServerNEXProtocols() {
	_ = utility.NewProtocol(globals.SecureServer)

}
