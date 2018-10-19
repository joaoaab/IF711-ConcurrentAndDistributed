package com.example.joaoaab.roommiddleware;

import android.os.StrictMode;
import android.support.v7.app.AppCompatActivity;
import android.os.Bundle;
import android.view.View;
import android.widget.Button;
import android.widget.TextView;

public class MainActivity extends AppCompatActivity {
    private Proxy board;

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_main);
        StrictMode.ThreadPolicy policy = new StrictMode.ThreadPolicy.Builder().permitAll().build();
        StrictMode.setThreadPolicy(policy);
        board = new Proxy();
        Button btnLed = (Button) findViewById(R.id.changeLedBtn);
        btnLed.setOnClickListener(new View.OnClickListener(){
            @Override
            public void onClick(View v){
                int ledStatus = board.turnLed();
                TextView test = findViewById(R.id.timeText);
                test.setText(Integer.toString(ledStatus));
            }
        });
    }


}
