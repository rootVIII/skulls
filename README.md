### Skulls is simple Columns-like strategy game developed in Golang with the Ebiten library (for Android)


<img src="https://images2.imgbox.com/a6/ab/4hlQKK3q_o.png" alt="ex2"/>


###### gomobile, build .apk for development and testing:

<pre>
  <code>
// 1.
// Navigate to skulls/ and generate a <code>.apk</code> with skullsgomobile/:
gomobile build -target=android github.com/rootVIII/skulls/skullsgomobile


// 2.
// Install the newly created .apk into an already running Android Emulator (from Android Studio):
adb -s emulator-5554  install skullsgomobile.apk


// 3. 
// view logging output from the game:
adb logcat


// Note that I use a pixel4 emulator. I have an alias stored in my profile to open it easily via terminal with the command pixel4:
alias pixel4='$ANDROID_HOME/emulator/emulator -avd "Pixel_4_API_30"'
  </code>
</pre>


###### ebitenmobile, build .aar for Android Studio binding:

<pre>
  <code>
// 1.
// Navigate to skulls/ and generate the <code>.aar</code> binding:
ebitenmobile bind -target android -javapkg com.&lt;your-username&gt;.skulls -o skulls.aar github.com/rootVIII/skulls/skullsebitenbind


// 2.
// Open an Empty Activity in Android Studio and name it SkullsMobile


// 3.
// Import the new .aar as a module:
// Select File, New, New Module, Import .jar/.aar Package, select the previously built .aar named skulls.aar
// In app/build.gradle, add this line to the dependencies: compile project(':skulls')
dependencies {
    implementation 'androidx.appcompat:appcompat:1.3.0'
    implementation 'com.google.android.material:material:1.3.0'
    implementation 'androidx.constraintlayout:constraintlayout:2.0.4'
    testImplementation 'junit:junit:4.+'
    androidTestImplementation 'androidx.test.ext:junit:1.1.2'
    androidTestImplementation 'androidx.test.espresso:espresso-core:3.3.0'
    compile project(':skulls')
}
// Then synch the change to the build.gradle for the project


// 4.
// Place the following in app/src/main/java/&lt;your username&gt;/MainActivity.java:
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


// 5.
// Add a separate error handling class in app/src/main/java/&lt;your-username&gt;/EbitenViewWithErrorHandling.java
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


// 6.
// Add the below into app/src/main/res/AndroidManifest.xml:
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


// 7.
// The game should now load in one of the emulators that comes with Android Studio
  </code>
</pre>


<ul>
  <li>
    All development/debugging was done with the <b>gomobile</b> tool and <b>adb</b>.
  </li>
  <li>
    Font used for text: <a href="https://www.dafont.com/radioland.font">RADIOLAND.TTF</a> 
  </li>
  <li>
  All assets (images, audio, and font) were converted to <code>[]byte</code> using <a href="https://github.com/hajimehoshi/file2byteslice">file2byteslice</a>
  </li>
</ul>


This was developed on macOS Big Sur.
<hr>
<b>Author: rootVIII  2021</b>
<br><br>
