package goc

import (
	"bytes"
	"os"
	"testing"

	"gitlab.com/TimSatke/gog"

	"github.com/stretchr/testify/require"
)

func TestPrefsGetSet(t *testing.T) {
	require := require.New(t)

	os.Remove("define.yaml")       // delete the preferences to avoid that old properties are being loaded
	defer os.Remove("define.yaml") // also clean up after test

	// start testing routine
	buf := &bytes.Buffer{}
	Out = buf

	prefsSet("myTestKey1", "myTestValue1")
	prefsSet("myTestKey2", "myTestValue2")
	prefsSet("myTest.Key.3", "myTest.Value.3")
	prefsSet("a b c d", "some value")
	prefsSet(`l00K at my special characters: !@#$%^&*()_[];'\{}:"|`, "another value... I guess")

	require.NoError(prefsGet("myTestKey1"))
	require.NoError(prefsGet("myTestKey2"))
	require.NoError(prefsGet("myTest.Key.3"))
	require.NoError(prefsGet("a b c d"))
	require.NoError(prefsGet(`l00K at my special characters: !@#$%^&*()_[];'\{}:"|`))
	require.NoError(prefsGet("Does not exist"))
	require.NoError(prefsGet("does not yet exist"))

	prefsSet("does not yet exist", "but now it does")
	require.NoError(prefsGet("does not yet exist"))

	prefsSet("myTest.Key.3", "myTest.Value.4")
	require.NoError(prefsGet("myTest.Key.3"))

	recorded := buf.Bytes()
	expected := gog.LoadGolden("TestPrefsGetSet", recorded)
	require.Equal(expected, recorded)
}
