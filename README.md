### Skulls! A simple Columns-like strategy game developed in Golang with the Ebiten library (for Android)

<a href="https://play.google.com/store/apps/details?id=com.solsticenet.skullsmobile" target="_blank">
  <img src="https://images2.imgbox.com/82/88/wEAnPcV2_o.png" alt="Google Play Store"/>
</a>
<hr>
<img src="https://images2.imgbox.com/a6/ab/4hlQKK3q_o.png" alt="ex1"/>
<hr>
<img src="https://images2.imgbox.com/5f/91/zXqDD7WR_o.png" alt="ex2"/>
<hr>



<ul>
  <li>
    The game was developed as a POC to experience creating a simple game with Go/deploying it to Android
  </li>
  <li>
    The <a href="https://ebiten.org/" target="_blank">Ebiten</a> library for Golang was used to create the game
  </li>
  <li>
    <a href="https://github.com/hajimehoshi/go-inovation" target="_blank">go-inovation</a> was used as a guide for the ebitenmobile .aar binding
  </li>
  <li>
    All development/debugging was done with the <a href="https://pkg.go.dev/golang.org/x/mobile/cmd/gomobile" target="_blank">gomobile</a> tool and <a href="https://developer.android.com/studio/command-line/adb" target="_blank">adb</a>
  </li>
  <li>
    Android Studio should be downloaded/installed; the AVD emulators are free and convenient
  </li>
  <li>
    I use the AVD emulators that are installable with Android Studio and stored in<br><code>$ANDROID_HOME/emulator/emulator</code> 
  </li>
  <li>
    It may be helpful to store an alias in your profile to open an emulator via a simple command:<br><code>alias pixel4='$ANDROID_HOME/emulator/emulator -avd "Pixel_4_API_30"'</code>
  </li>
  <li>
    Font used for text: <a href="https://www.dafont.com/radioland.font">RADIOLAND.TTF</a> 
  </li>
  <li>
    All assets/ (images, audio, and font) were converted to <code>[]byte</code> using <a href="https://github.com/hajimehoshi/file2byteslice">file2byteslice</a>
  </li>
  <li>
    The project is intended to be built with gomobile for development and testing, or with ebitenmobile for production releases using Android Studio
  </li>
</ul>

###### Build .apk for development and testing using gomobile:

<pre>
  <code>
// 1. Navigate to skulls/ and generate a <code>.apk</code> with gomobile:
gomobile build -target=android github.com/rootVIII/skulls/skullsgomobile



// 2. Install the newly created .apk into an already running Android Emulator:
adb -s &lt;emulator-name&gt; install skullsgomobile.apk

// Note: to list available emulators (including phone connected for debugging):
adb devices -l



// 3. View debug/logging output from the game:
adb logcat
  </code>
</pre>
<br>

###### Build .aar for Android Studio binding and production release using ebitenmobile:

<pre>
  <code>
// 1. Navigate to skulls/ and generate the <code>.aar</code> binding:
ebitenmobile bind -target android -javapkg com.&lt;your-username&gt;.skulls -o skulls.aar github.com/rootVIII/skulls/skullsebitenbind



// 2. Create a new Android Studio project (choose Empty Activity) and name it SkullsMobile



// 3. Import the new .aar as a module:
// Select File, New, New Module, Import .jar/.aar Package, select the previously built .aar named skulls.aar
// In app/build.gradle, add this line to the dependencies: compile project(':skulls')
// Example:

dependencies {
    implementation 'androidx.appcompat:appcompat:1.3.0'
    implementation 'com.google.android.material:material:1.3.0'
    implementation 'androidx.constraintlayout:constraintlayout:2.0.4'
    testImplementation 'junit:junit:4.+'
    androidTestImplementation 'androidx.test.ext:junit:1.1.2'
    androidTestImplementation 'androidx.test.espresso:espresso-core:3.3.0'
    compile project(':skulls')
}
// Then follow screen prompts to sync the build.gradle change to the project



// 4. Place the following in app/src/main/java/com.&lt;your username&gt;.skullsmobile/MainActivity.java:

package com.&lt;your-username&gt;.skullsmobile;

import androidx.appcompat.app.AppCompatActivity;

import android.os.Bundle;

import go.Seq;
import com.&lt;your-username&gt;.skulls.skullsebitenbind.EbitenView;


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



// 5. Add a separate error handling class in app/src/main/java/com.&lt;your-username&gt;skullsmobile/EbitenViewWithErrorHandling.java

package com.solsticenet.skullsmobile;

import android.content.Context;
import android.util.AttributeSet;

import com.&lt;your-username&gt;.skulls.skullsebitenbind.EbitenView;


class EbitenViewWithErrorHandling extends EbitenView {
    public EbitenViewWithErrorHandling(Context context) {
        super(context);
    }

    public EbitenViewWithErrorHandling(Context context, AttributeSet attributeSet) {
        super(context, attributeSet);
    }

    @Override
    protected void onErrorOnGameUpdate(Exception e) {
        // You can define your own error handling e.g., using Crashlytics.
        // e.g., Crashlytics.logException(e);
        super.onErrorOnGameUpdate(e);
    }
}



// 6. Add the below into app/src/main/res/AndroidManifest.xml:
&lt;?xml version="1.0" encoding="utf-8"?&gt;
&lt;RelativeLayout xmlns:android="http://schemas.android.com/apk/res/android"
    xmlns:tools="http://schemas.android.com/tools"
    android:layout_width="match_parent"
    android:layout_height="match_parent"
    android:background="@color/background_material_dark"
    android:keepScreenOn="true"
    android:screenOrientation="portrait"
    tools:context="com.&lt;your-username&gt;.skullsmobile.MainActivity"&gt;

    &lt;com.&lt;your-username&gt;.skullsmobile.EbitenViewWithErrorHandling
        android:id="@+id/ebitenview"
        android:layout_width="match_parent"
        android:layout_height="match_parent"
        android:focusable="true" /&gt;
&lt;/RelativeLayout&gt;



// 7. The game should now be usable in Android Studio (sign the project with developer keys, UI adjustments in AndroidManifest.xml etc.)
  </code>
</pre>

<br>

This was developed on macOS Big Sur.
<hr>
<b>Author: rootVIII  2021</b>
<br><br>
