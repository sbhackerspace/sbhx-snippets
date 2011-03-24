
(defn test[]
  (do
    (def five 5)
    (println (* 5 5))
    (println 'five)))

(defn hello-you [name]
  (println "Hello" name ))

;(test)
;(hello-you "Billy The Kid")

(let [[bob & the-gang :as everyone]["bob" "kizzy" "gabe" "john" "beans"]]
(list bob the-gang everyone))