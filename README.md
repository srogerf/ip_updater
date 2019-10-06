# ip_updater


./ip_updater -domain=${DOMAIN} -key=${API_KEY} -secret=${API_SECRET}


https://developer.godaddy.com/getstarted

curl -X GET -H"Authorization: sso-key ${API_KEY}:${API_SECRET}" "https://api.godaddy.com/v1/domains/available?domain=example.guru"
