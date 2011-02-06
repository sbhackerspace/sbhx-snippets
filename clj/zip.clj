; Steve Phillips / elimisteve
; 2011.01.26

(require '(clojure [zip :as zip]))

(def my-data [[[1 4] [7 9]] [2 8]])
(def my-vzip (zip/vector-zip my-data))

(-> my-vzip zip/node)
; [[[1 4] [7 9]] [2 8]]

;; Are these always equivalent?
(-> my-vzip zip/next zip/node)
(-> my-vzip zip/down zip/node)
; [[1 4] [7 9]]

(-> my-vzip zip/next zip/right zip/node)
; [2 8]

(-> my-vzip zip/next zip/right zip/down zip/node)
; 2

(-> my-vzip zip/next zip/right zip/down zip/right zip/node)
; 8


;;;; Creates namespace conflicts, but more convenient

(use 'clojure.zip)
(def my-data [[[1 4] [7 9]] [2 8]])
(def my-vzip (vector-zip my-data))

(-> my-vzip node)
; [[[1 4] [7 9]] [2 8]]

;; Are these always equivalent?
(-> my-vzip next node)
(-> my-vzip down node)
; [[1 4] [7 9]]

(-> my-vzip next right node)
; [2 8]
;; Without ->, this would look like
;(node (right (next my-vzip)))


(-> my-vzip next right down node)
; 2

(-> my-vzip next right down right node)
