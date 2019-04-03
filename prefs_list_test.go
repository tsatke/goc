package goc

import (
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/TimSatke/gog"
)

func TestPrefsList(t *testing.T) {
	require := require.New(t)

	os.Remove("define.yaml")       // delete the preferences to avoid that old properties are being loaded
	defer os.Remove("define.yaml") // also clean up after test

	buf := &bytes.Buffer{}
	Out = buf

	prefsSet("cmd.output.directory", "/var/lib/tools")
	prefsSet("myTestKey1", "myTestValue1")
	prefsSet("myTestKey2", "myTestValue2")
	prefsSet("myTest.Key.3", "myTest.Value.3")
	prefsSet("a b c d", "some value")
	prefsSet(`l00K at my special characters: !@#$%^&*()_[];'\{}:"|`, "another value... I guess")

	prefsList()

	recorded := buf.Bytes()
	expected := gog.LoadGolden("TestPrefsList", recorded)
	require.Equal(expected, recorded)
}
