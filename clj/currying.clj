;;William Duncan, credit due to the author of Practical Clojure.

;;cool, you can create new functions by closing over some of the params.
(def times-two (partial * 2))

;;returns function that performs func2 on the params, then func1 ond the result
(comp func1 func2)

;;defines fn that multiplies two params, negates the result, the mult by ten. Good example of using comp and partial together in currying.
(def my-fn (comp (partial * 10) - *))

