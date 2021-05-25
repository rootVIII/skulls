### Skulls is simple Columns-like strategy game developed in Golang with the Ebiten library (for Android)


All development/debugging was done with the <b>gomobile</b> tool and <b>adb</b>.


###### DEVELOPMENT: Local testing environment

<pre>
  <code>
// Navigate to skulls/ and generate a <code>.apk</code> with skullsgomobile/:
gomobile build -target=android github.com/rootVIII/skulls/skullsgomobile

// Install the newly created .apk into an already running Android Emulator (from Android Studio):
adb -s emulator-5554  install skullsgomobile.apk

// view logging output from the game:
adb logcat

// Note that I use a pixel4 emulator. I have an alias stored in my profile to open it easily via terminal:
alias pixel4='$ANDROID_HOME/emulator/emulator -avd "Pixel_4_API_30"'
  </code>
</pre>

###### PRODUCTION: Build for Android Studio

<pre>
  <code>
// Navigate to skulls/ and generate a .aar with skullsebitenbind/
ebitenmobile bind -target android -javapkg com.solsticenet.skulls -o skulls.aar github.com/rootVIII/skulls/skullsebitenbind
  
// Create a new project in Android Studio named SkullsMobile.
// Click File, New Module, import .jar/.aar, and locate the newly created skulls.aar file.
// Add the following line into dependencies{} within app/build.gradle and then clean the project:
compile project(':skulls')

// Place the following code inside of MainActivity.java:
package com.solsticenet.skullsmobile;

import androidx.appcompat.app.AppCompatActivity;

import android.os.Bundle;

import go.Seq;
import com.solsticenet.skulls.skullsebitenbind.EbitenView;

public class MainActivity extends AppCompatActivity {

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_main);
        Seq.setContext(getApplicationContext());
    }

    private EbitenView getEbitenView() {
        return (EbitenView)this.findViewById(R.id.ebitenview);
    }

    @Override
    protected void onPause() {
        super.onPause();
        this.getEbitenView().suspendGame();
    }

    @Override
    protected void onResume() {
        super.onResume();
        this.getEbitenView().resumeGame();
    }
}
  </code>
</pre>


This was developed on macOS Big Sur.
<hr>
<b>Author: rootVIII  2021</b>
<br><br>
