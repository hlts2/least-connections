package leastconnections

// func TestLeastConnections(t *testing.T) {
// 	lc, err := New([]string{
// 		"server-1",
// 		"server-2",
// 		"server-3",
// 	})
//
// 	if err != nil {
// 		t.Error("New is nil")
// 	}
//
// 	server1 := lc.Next()
// 	if "server-1" != server1 {
// 		t.Errorf("Next is not wrong. expected: %v, got: %v", "server-1", server1)
// 	}
// 	lc.IncrementConnections(server1)
//
// 	server2 := lc.Next()
// 	if "server-2" != server2 {
// 		t.Errorf("Next is not wrong. expected: %v, got: %v", "server-2", server2)
// 	}
// 	lc.IncrementConnections(server2)
//
// 	lc.DecrementConnections(server1)
//
// 	server1 = lc.Next()
// 	if "server-1" != server1 {
// 		t.Errorf("Next is not wrong. expected: %v, got: %v", "server-1", server1)
// 	}
//
// 	lc.IncrementConnections(server1)
//
// 	server3 := lc.Next()
// 	if "server-3" != server3 {
// 		t.Errorf("Next is not wrong. expected: %v, got: %v", "server-3", server3)
// 	}
//
// 	lc.DecrementConnections(server1)
// 	lc.DecrementConnections(server2)
// 	lc.DecrementConnections(server3)
//
// 	server1 = lc.Next()
// 	if "server-1" != server1 {
// 		t.Errorf("Next is not wrong. expected: %v, got: %v", "server-1", server1)
// 	}
// }
