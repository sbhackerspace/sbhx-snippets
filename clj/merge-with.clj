; Steve Phillips / elimisteve
; 2011.03.17

(def big-string "The map get function (get) takes a map and a key as its
first and second arguments, and an optional third argument specifying
the value if the key is not found. It returns the value of the
specified key in the map, returning nil if it is not found and there
is no third argument.")

(apply merge-with +
	   (for [letter big-string]
		 {letter 1}))
