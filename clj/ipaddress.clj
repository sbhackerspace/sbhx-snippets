; From http://ktuman.blogspot.com/2011/01/detect-your-host-ip-with-clojure.html

(import (java.net NetworkInterface))

(def ip
	 (let [ifc (NetworkInterface/getNetworkInterfaces)
		   ifsq (enumeration-seq ifc)
		   ifmp (map #(bean %) ifsq)
		   ipsq (filter #(false? (% :loopback)) ifmp)
		   ipa (map :interfaceAddresses ipsq)
		   ipaf (nth ipa 0)
		   ipafs (.split (str ipaf) " " )
		   ips (first (nnext ipafs))]
	   (str (second (.split ips "/")))))

(println ip)
