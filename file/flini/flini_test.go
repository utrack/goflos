package flini_test

import (
	"testing"

	"strings"

	"github.com/stretchr/testify/assert"
	"github.com/utrack/goflos/file/flini"
)

func TestIni(t *testing.T) {
	so := assert.New(t)

	ini, err := flini.Parse(strings.NewReader(testdata))
	so.NoError(err)

	so.Equal(len(ini.Sections["Freelancer"].Values), 4)
	so.Equal(ini.Sections["Freelancer"].Values["initial_world"][0], "InitialWorld.ini")
}

var testdata = `
[Freelancer]
data path= ..\data
local_server = rpclocal.dll
initial_world = InitialWorld.ini	;relative to Data path
AppGUID = {A690F026-26F0-4e57-ACA0-ECF868E48D21}
    ; some uneven comment
; some comment


[;Display]
    fullscreen = 1
    size = 1024,768
    color_bpp = 32
    depth_bpp = 32

[Startup]
movie_file = movies\MGS_Logo_Final.wmv
movie_file = movies\DA_Logo_Final.wmv
movie_file = movies\FL_Intro.wmv


[ListServer]
;;;hostname = localhost                              ;Your local machine
;;;hostname = FLListServer2.dns.corp.microsoft.com   ;GUN server in Austin
;;;hostname = 131.107.135.190                        ;GUN server in Redmond
hostname = fllistserver.zone.msn.com              ;GUN server in Redmond (DNS entry)

port = 2300

[Server]
;name = M9Universe
;description = My cool Freelancer server
death_penalty = 100   ; percentage of your cargo (commoditied and unmounted equipment) lost at death in MP

[Initial MP DLLs]
path = ..\dlls\bin
DLL = Content.dll, GameSupport, HIGHEST ; test comment
; required to operate gates and docks
; required to create ships in space
DLL = Content.dll, SpaceSupport, NORMAL 
DLL = Content.dll, BaseSupport, NORMAL

DLL = Content.dll, SpacePop, LOWEST ;populator
DLL = Content.dll, AISandbox, BELOW_NORMAL
DLL = Content.dll, TestAutomation, BELOW_NORMAL
DLL = Content.dll, BasePop, LOWEST
`
