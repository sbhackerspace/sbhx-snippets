; Steve Phillips / elimisteve
; 2011.01.26
(max-key #(nth % 0) [1 2 3] [6 0 3])
; [6 0 3]

(max-key (fn [x] (nth x 0)) [1 2 3] [6 0 3])
; [6 0 3]

(max-key #(nth % 1) [1 2 3] [6 0 3])
; [1 2 3]

(max-key #(nth % 2) [1 2 3] [6 0 3])
; [6 0 3]


;; Using 'first' or 'second' is cleaner when you can use it
(min-key first [1 7] [3 5] [4 4])
; [1 7]

(max-key first [1 7] [3 5] [4 4])
; [4 4]


;; But what do you do when your input is a list?
(min-key first (list [1 7] [3 5] [4 4]))
; ([1 7] [3 5] [4 4])  <-- Not what we want

;; Answer: use 'apply'
(apply min-key first (list [1 7] [3 5] [4 4]))
; [1 7]

(apply min-key first '([1 7] [3 5] [4 4]))
; [1 7]


;; From http://www.cs.ucla.edu/~rosen/161/notes/alphabeta.html
;; converted to (mostly) vectors
(def my-tree
	 '([[[3 17] [2 12]]
		[[15] [25 0]]]
		 [[[2 5] [3]]
		  [[2 14]]]))

;; (defn print-first-leaf [atree]
;;   (if-not (integer? atree)
;; 	(print-nodes (first atree))
;; 	atree))

;; (print-first-leaf my-tree)

;; Also prints first leaf
(first (filter integer? (iterate first my-tree)))

(first (first (first (first my-tree))))        ; 3
(second (first (first (first my-tree))))       ; 17
(second (first (first (second my-tree))))      ; 5
(second (second (first (first my-tree))))      ; 12
(first (first (second (first my-tree))))       ; 15
(second (first (second (second my-tree))))     ; 14

(-> my-tree second first first second)

;; (if 
;; (for [elt my-tree]
;;   elt)
