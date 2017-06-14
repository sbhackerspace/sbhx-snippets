;; Steve Phillips / elimisteve
;; 2017.06.14

(require '[org.httpkit.client :as http])

(defn- unescape-html
  [s]
  (-> s (.replace "&amp;" "&") (.replace "&lt;" "<") (.replace "&gt;" ">")))

(let [{:keys [status headers body error] :as resp} @(http/get "https://clojure.github.io/clojure/clojure.core-api.html")]
  (if error
    (println "Failed, exception: " error)
    (let [h2s (map (comp unescape-html second)
                   (re-seq #"<h2 id=\"(?:.*?)\">(.*?)</h2>" body))]
      (printf "There are %d keywords in the latest version of Clojure: %s"
              (count h2s)
              (string/join " " h2s)))))
