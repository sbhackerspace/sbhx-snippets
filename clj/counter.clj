;;William Duncan 3/26/11

(defn counter [x]
    (fn []
      (def x (inc x))
      x))

;;why doesn't this work?