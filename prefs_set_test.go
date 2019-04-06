package goc

import (
	"bytes"
	"os"
	"testing"

	"gitlab.com/TimSatke/gog"

	"github.com/stretchr/testify/require"
)

func TestPrefsGetSet(t *testing.T) {
	const (
		testKey1 = "myTestKey1"
		testKey2 = "myTestKey2"
		testKey3 = "myTest.Key.3"
		testKey4 = "a b c d"
		testKey5 = `l00K at my special characters: !@#$%^&*()_[];'\{}:"|`
		testKey6 = "Does not exist"
		testKey7 = "does not yet exist"
	)

	require := require.New(t)

	os.Remove("define.yaml")       // delete the preferences to avoid that old properties are being loaded
	defer os.Remove("define.yaml") // also clean up after test

	// start testing routine
	buf := &bytes.Buffer{}
	Out = buf

	prefsSet(testKey1, "myTestValue1")
	prefsSet(testKey2, "myTestValue2")
	prefsSet(testKey3, "myTest.Value.3")
	prefsSet(testKey4, "some value")
	prefsSet(testKey5, "another value... I guess")

	require.NoError(prefsGet(testKey1))
	require.NoError(prefsGet(testKey2))
	require.NoError(prefsGet(testKey3))
	require.NoError(prefsGet(testKey4))
	require.NoError(prefsGet(testKey5))
	require.NoError(prefsGet(testKey6))
	require.NoError(prefsGet(testKey7))

	prefsSet(testKey7, "but now it does")
	require.NoError(prefsGet(testKey7))

	prefsSet(testKey3, "myTest.Value.4")
	require.NoError(prefsGet(testKey3))

	recorded := buf.Bytes()
	expected := gog.LoadGolden("TestPrefsGetSet", recorded)
	require.Equal(expected, recorded)
}
