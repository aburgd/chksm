package main

import (
	"encoding/hex"
	"reflect"
	"testing"
)

func TestSha256Sum(t *testing.T) {
	// get the sha256 for a, then make a string out of it
	testHashA := sha256Sum("test-file-a")
	testHashHexA := hex.EncodeToString(testHashA)

	// ditto for b
	testHashB := sha256Sum("test-file-b")
	testHashHexB := hex.EncodeToString(testHashB)

	// compare the sha256 from the source with sha256 obtained from
	// other sources
	assertHashA := "4661aeaf9f4a9a7f5644b6af88623b66a35a494aa06177ed358adfd6b5877207"
	assertHashB := "97ba5787259a6798aec95c2df2ef597c6839fb71598bdabbd71d9cc60dc95047"

	equalityA := reflect.DeepEqual(assertHashA, testHashHexA)
	equalityB := reflect.DeepEqual(assertHashB, testHashHexB)

	if !equalityA {
		t.Errorf("SHA256 for test-file-a was incorrect, wanted %s, got %s", assertHashA, testHashHexA)
	}

	if !equalityB {
		t.Errorf("SHA256 for test-file-a was incorrect, wanted %s, got %s", assertHashB, testHashHexB)
	}

}

func TestSha1Sum(t *testing.T) {
	testHashA := sha1Sum("test-file-a")
	testHashHexA := hex.EncodeToString(testHashA)

	testHashB := sha1Sum("test-file-b")
	testHashHexB := hex.EncodeToString(testHashB)

	assertHashA := "2a428c0c963b2ff59d89d521a8bd3d085a42dd92"
	assertHashB := "56e7ca5b1290b54f2ae793a41d4b0382a1a3b515"

	equalityA := reflect.DeepEqual(assertHashA, testHashHexA)
	equalityB := reflect.DeepEqual(assertHashB, testHashHexB)

	if !equalityA {
		t.Errorf("SHA1 for test-file-a was incorrect, wanted %s, got %s", assertHashA, testHashHexA)
	}

	if !equalityB {
		t.Errorf("SHA1 for test-file-b was incorrect, wanted %s, got %s", assertHashB, testHashHexB)
	}

}

func TestMd5Sum(t *testing.T) {
	testHashA := md5Sum("test-file-a")
	testHashHexA := hex.EncodeToString(testHashA)
	testHashB := md5Sum("test-file-b")
	testHashHexB := hex.EncodeToString(testHashB)

	assertHashA := "ec5f8b842a5c60c8fb05a884e5a33b71"
	assertHashB := "3ca34b22875efdfa647f03e8f7aa5ea3"

	equalityA := reflect.DeepEqual(assertHashA, testHashHexA)
	equalityB := reflect.DeepEqual(assertHashB, testHashHexB)

	if !equalityA {
		t.Errorf("MD5 for test-file-a was incorrect, wanted %s, got %s", assertHashA, testHashHexA)
	}

	if !equalityB {
		t.Errorf("MD5 for test-file-b was incorrect, wanted %s, got %s", assertHashB, testHashHexB)
	}

}
