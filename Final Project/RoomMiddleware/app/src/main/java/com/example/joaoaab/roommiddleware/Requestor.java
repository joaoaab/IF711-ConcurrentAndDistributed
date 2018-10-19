package com.example.joaoaab.roommiddleware;
import android.util.Log;

import org.json.JSONException;
import org.json.JSONObject;


public class Requestor {
    private ClientRequestHandler crh;

    public Requestor(){
        this.crh = new ClientRequestHandler("192.168.100.3", 700);
    }

    public int Invoke(int device, int op, int value){
        JSONObject obj = new JSONObject();

        try{
            obj.put("device", device);
            obj.put("operation", op);
            obj.put("value", value);
        }catch(JSONException e){
            e.printStackTrace();
        }

        String msg = "";
        msg = obj.toString();
        Log.d("Tomele", msg);
        this.crh.send(msg);


        try{
            Thread.sleep(800);
        }catch(InterruptedException e){
            e.printStackTrace();
        }

        String ans = "";
        ans = this.crh.receive();
        JSONObject ret = new JSONObject();
        try{
            ret = new JSONObject(ans);
        }catch(JSONException e){
            e.printStackTrace();
        }
        int val = -1;
        try{
            val = ret.getInt("value");
        }catch(JSONException e){
            e.printStackTrace();

        }
        return val;
    }



}
