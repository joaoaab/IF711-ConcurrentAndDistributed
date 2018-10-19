package com.example.joaoaab.roommiddleware;

import android.os.AsyncTask;
import android.util.Log;

import java.io.BufferedReader;
import java.io.BufferedWriter;
import java.io.DataInputStream;
import java.io.DataOutputStream;
import java.io.IOException;
import java.io.UnsupportedEncodingException;
import java.net.Socket;

public class ClientRequestHandler {
    private String host;
    private int port;
    private int sentMessageSize;
    private int receiveMessageSize;

    private Socket clientSocket = null;
    private DataOutputStream outToServer = null;
    private DataInputStream inFromServer = null;


    public ClientRequestHandler(String host, int port){
        this.host = host;
        this.port = port;
        try{
            this.clientSocket = new Socket(host, port);
            this.outToServer = new DataOutputStream(clientSocket.getOutputStream());
            this.inFromServer = new DataInputStream(clientSocket.getInputStream());
        }catch(IOException e) {
            e.printStackTrace();
        }
    }

    public void send(String msg){
        Log.d("send", msg);
        try{
            this.outToServer.writeBytes(msg);
        }catch(IOException e){
            e.printStackTrace();
        }
    }

    public String receive(){
        byte []msg = new byte[1024];
        String decoded = "";
        int read = -1;
        try{
            read = this.inFromServer.read(msg);
            Log.d("bytesRead", Integer.toString(read));
        }catch (IOException e) {
            e.printStackTrace();
        }
        if(read != -1){
            decoded = new String(msg, 0, read);
        }
        Log.d("received", decoded);
        return decoded;
    }


}
