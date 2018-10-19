package com.example.joaoaab.roommiddleware;

public class Proxy {
    private Requestor requestor;

    public Proxy(){
        this.requestor = new Requestor();
    }

    public int turnLed(){
        return this.requestor.Invoke(0, 0, 1);
    }
}
