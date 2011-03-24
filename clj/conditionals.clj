;;William Duncan 3/23/11
;; if-not is a negated if statement
(if-not true
  (println "true is false")
  (println "true is true"))

;;cond can be given an :else keyword for a default behavior
(cond
 (= 2 1) "What?"
 (= 5 4) "What What?"
 :else "Ah, made it out of the storm")

