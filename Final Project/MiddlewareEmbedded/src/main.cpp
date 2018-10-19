#include <ESP8266WiFi.h>
#include <ArduinoJson.h>

#define LED 1
#define LDR 2
#define TOMELE 3

char ssid[] = "SSID";
char password[] = "123456";
int ledStatus = -1;
StaticJsonBuffer<512> bufferIn;
StaticJsonBuffer<512> bufferOut;

WiFiServer server(700);

void setup() {
    Serial.begin(9600);
    while(!Serial){
    }


    Serial.println("Trying to connect to network");
    Serial.println(ssid);

    int status = WiFi.begin(ssid, password);
    status = WiFi.waitForConnectResult();

    if(status != WL_CONNECTED){
        Serial.println("Failed to connect");
        while(true){
            Serial.println("kk");
        }
    }
    Serial.println("Connected.");
    Serial.print("MAC Addr: ");
    Serial.println(WiFi.macAddress());
    Serial.print("IP Addr:  ");
    Serial.println(WiFi.localIP());
    Serial.print("Subnet:   ");
    Serial.println(WiFi.subnetMask());
    Serial.print("Gateway:  ");
    Serial.println(WiFi.gatewayIP());
    Serial.print("DNS Addr: ");
    Serial.println(WiFi.dnsIP());
    Serial.print("Channel:  ");
    Serial.println(WiFi.channel());
    Serial.print("Status: ");
    Serial.println(WiFi.status());

    server.begin();
}


int processMsg(WiFiClient client){
    JsonObject &root = bufferIn.parse(client);
    int device = root["device"];
    int operation = root["operation"];
    int value = root["value"];
    if(device == 0){
        if(operation == 0){
            ledStatus = value;
            return ledStatus;
        }
    }
    return -1;
}

void answer(int ret, WiFiClient client){
    JsonObject &root = bufferOut.createObject();
    root["value"] = ret;
    root.printTo(client);
}


void loop() {
    WiFiClient client = server.available();
    if(client){
        while(client.connected()){
            if(client.available()){
                int ret = processMsg(client);
                answer(ret, client);
            }
        }
    }
    // Change Led Status
}