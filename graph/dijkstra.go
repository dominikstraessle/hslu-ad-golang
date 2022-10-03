package graph

//func shortPathDijkstra(s int, z int) int {
//	// Einfach implementiert,
//	var noOfNodes = 0
//	var adjaMxb [1][1]int
//	var setT = [noOfNodes]bool{}
//	// nicht optimiert!
//	for i := 0; i < noOfNodes; i++ {
//		setT[i] = false
//	}
//	setT[s] = true
//	// Startknoten s ist Element von T.
//	var minEntf [noOfNodes]int
//	minEntf[s] = 0
//	// K?rzeste Entfernung zum Startknoten ist null.
//	for !setT[z] {
//		// Terminiert, wenn Ziel z erreicht.
//		var min int = math.MaxInt64
//		// K?rzeste Entfernung von k zum n?chsten nk.
//		var kNeu int = 0
//		// Nachbar nk mit k?rzester Entfernung zu k.
//		for k := 0; k < noOfNodes; k++ {
//			if setT[k] {
//				// Falls k in Menge T, ...
//				for nk := 0; nk < noOfNodes; nk++ {
//					// Falls Nachbar nk, ...
//					if !setT[nk] && (adjaMxb[k][nk] < math.MaxInt64) {
//						var kNeuEntf int = minEntf[k] + adjaMxb[k][nk]
//						if kNeuEntf < min {
//							min = kNeuEntf
//							kNeu = nk
//						}
//					}
//				}
//			}
//		}
//		setT[kNeu] = true
//		// kNeu wird in T aufgenommen
//		minEntf[kNeu] = min
//	}
//	return minEntf[z]
//}
