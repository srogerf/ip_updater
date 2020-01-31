# ip_updater


./ip_updater -domain=${DOMAIN} -key=${API_KEY} -secret=${API_SECRET}


###Godaddy API documentation
https://developer.godaddy.com/doc/endpoint/domains

###Godaddy api_key, api_secret setup
https://developer.godaddy.com/getstarted

curl -X GET -H"Authorization: sso-key ${API_KEY}:${API_SECRET}" "https://api.godaddy.com/v1/domains/available?domain=example.guru"


# UI
react

cd frontend/dns-manager
npm run build
cd ../..
./ip_updater <options> -daemon


dev
cd frontend/dns-manager
npm start

