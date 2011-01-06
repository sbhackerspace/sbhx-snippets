(defn many-ifs [x]
  (cond
   (< x 0) (println "Negative")
   (= x 0) (println "Zero")
   (> x 0) (println "Positive")))

(many-ifs 5)