// Code generated by SQLBoiler 4.6.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import "testing"

// This test suite runs each operation test in parallel.
// Example, if your database has 3 tables, the suite will run:
// table1, table2 and table3 Delete in parallel
// table1, table2 and table3 Insert in parallel, and so forth.
// It does NOT run each operation group in parallel.
// Separating the tests thusly grants avoidance of Postgres deadlocks.
func TestParent(t *testing.T) {
	t.Run("Crawls", testCrawls)
	t.Run("PeerProperties", testPeerProperties)
	t.Run("Peers", testPeers)
	t.Run("Sessions", testSessions)
	t.Run("Tests", testTests)
}

func TestDelete(t *testing.T) {
	t.Run("Crawls", testCrawlsDelete)
	t.Run("PeerProperties", testPeerPropertiesDelete)
	t.Run("Peers", testPeersDelete)
	t.Run("Sessions", testSessionsDelete)
	t.Run("Tests", testTestsDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("Crawls", testCrawlsQueryDeleteAll)
	t.Run("PeerProperties", testPeerPropertiesQueryDeleteAll)
	t.Run("Peers", testPeersQueryDeleteAll)
	t.Run("Sessions", testSessionsQueryDeleteAll)
	t.Run("Tests", testTestsQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("Crawls", testCrawlsSliceDeleteAll)
	t.Run("PeerProperties", testPeerPropertiesSliceDeleteAll)
	t.Run("Peers", testPeersSliceDeleteAll)
	t.Run("Sessions", testSessionsSliceDeleteAll)
	t.Run("Tests", testTestsSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("Crawls", testCrawlsExists)
	t.Run("PeerProperties", testPeerPropertiesExists)
	t.Run("Peers", testPeersExists)
	t.Run("Sessions", testSessionsExists)
	t.Run("Tests", testTestsExists)
}

func TestFind(t *testing.T) {
	t.Run("Crawls", testCrawlsFind)
	t.Run("PeerProperties", testPeerPropertiesFind)
	t.Run("Peers", testPeersFind)
	t.Run("Sessions", testSessionsFind)
	t.Run("Tests", testTestsFind)
}

func TestBind(t *testing.T) {
	t.Run("Crawls", testCrawlsBind)
	t.Run("PeerProperties", testPeerPropertiesBind)
	t.Run("Peers", testPeersBind)
	t.Run("Sessions", testSessionsBind)
	t.Run("Tests", testTestsBind)
}

func TestOne(t *testing.T) {
	t.Run("Crawls", testCrawlsOne)
	t.Run("PeerProperties", testPeerPropertiesOne)
	t.Run("Peers", testPeersOne)
	t.Run("Sessions", testSessionsOne)
	t.Run("Tests", testTestsOne)
}

func TestAll(t *testing.T) {
	t.Run("Crawls", testCrawlsAll)
	t.Run("PeerProperties", testPeerPropertiesAll)
	t.Run("Peers", testPeersAll)
	t.Run("Sessions", testSessionsAll)
	t.Run("Tests", testTestsAll)
}

func TestCount(t *testing.T) {
	t.Run("Crawls", testCrawlsCount)
	t.Run("PeerProperties", testPeerPropertiesCount)
	t.Run("Peers", testPeersCount)
	t.Run("Sessions", testSessionsCount)
	t.Run("Tests", testTestsCount)
}

func TestHooks(t *testing.T) {
	t.Run("Crawls", testCrawlsHooks)
	t.Run("PeerProperties", testPeerPropertiesHooks)
	t.Run("Peers", testPeersHooks)
	t.Run("Sessions", testSessionsHooks)
	t.Run("Tests", testTestsHooks)
}

func TestInsert(t *testing.T) {
	t.Run("Crawls", testCrawlsInsert)
	t.Run("Crawls", testCrawlsInsertWhitelist)
	t.Run("PeerProperties", testPeerPropertiesInsert)
	t.Run("PeerProperties", testPeerPropertiesInsertWhitelist)
	t.Run("Peers", testPeersInsert)
	t.Run("Peers", testPeersInsertWhitelist)
	t.Run("Sessions", testSessionsInsert)
	t.Run("Sessions", testSessionsInsertWhitelist)
	t.Run("Tests", testTestsInsert)
	t.Run("Tests", testTestsInsertWhitelist)
}

// TestToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestToOne(t *testing.T) {
	t.Run("PeerPropertyToCrawlUsingCrawl", testPeerPropertyToOneCrawlUsingCrawl)
	t.Run("SessionToPeerUsingPeer", testSessionToOnePeerUsingPeer)
}

// TestOneToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOne(t *testing.T) {}

// TestToMany tests cannot be run in parallel
// or deadlocks can occur.
func TestToMany(t *testing.T) {
	t.Run("CrawlToPeerProperties", testCrawlToManyPeerProperties)
	t.Run("PeerToSessions", testPeerToManySessions)
}

// TestToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneSet(t *testing.T) {
	t.Run("PeerPropertyToCrawlUsingPeerProperties", testPeerPropertyToOneSetOpCrawlUsingCrawl)
	t.Run("SessionToPeerUsingSessions", testSessionToOneSetOpPeerUsingPeer)
}

// TestToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneRemove(t *testing.T) {}

// TestOneToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneSet(t *testing.T) {}

// TestOneToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneRemove(t *testing.T) {}

// TestToManyAdd tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyAdd(t *testing.T) {
	t.Run("CrawlToPeerProperties", testCrawlToManyAddOpPeerProperties)
	t.Run("PeerToSessions", testPeerToManyAddOpSessions)
}

// TestToManySet tests cannot be run in parallel
// or deadlocks can occur.
func TestToManySet(t *testing.T) {}

// TestToManyRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyRemove(t *testing.T) {}

func TestReload(t *testing.T) {
	t.Run("Crawls", testCrawlsReload)
	t.Run("PeerProperties", testPeerPropertiesReload)
	t.Run("Peers", testPeersReload)
	t.Run("Sessions", testSessionsReload)
	t.Run("Tests", testTestsReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("Crawls", testCrawlsReloadAll)
	t.Run("PeerProperties", testPeerPropertiesReloadAll)
	t.Run("Peers", testPeersReloadAll)
	t.Run("Sessions", testSessionsReloadAll)
	t.Run("Tests", testTestsReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("Crawls", testCrawlsSelect)
	t.Run("PeerProperties", testPeerPropertiesSelect)
	t.Run("Peers", testPeersSelect)
	t.Run("Sessions", testSessionsSelect)
	t.Run("Tests", testTestsSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("Crawls", testCrawlsUpdate)
	t.Run("PeerProperties", testPeerPropertiesUpdate)
	t.Run("Peers", testPeersUpdate)
	t.Run("Sessions", testSessionsUpdate)
	t.Run("Tests", testTestsUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("Crawls", testCrawlsSliceUpdateAll)
	t.Run("PeerProperties", testPeerPropertiesSliceUpdateAll)
	t.Run("Peers", testPeersSliceUpdateAll)
	t.Run("Sessions", testSessionsSliceUpdateAll)
	t.Run("Tests", testTestsSliceUpdateAll)
}
