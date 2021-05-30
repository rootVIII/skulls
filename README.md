### Skulls is simple Columns-like strategy game developed in Golang with the Ebiten library (for Android)


<img src="https://images2.imgbox.com/a6/ab/4hlQKK3q_o.png" alt="ex2"/>


###### gomobile, build .apk for development and testing:

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


###### ebitenmobile, build .aar for Android Studio binding:

<pre>
  <code>
// Navigate to skulls/ and generate the <code>.aar</code> binding:
ebitenmobile bind -target android -javapkg com.&lt;your-username&gt;.skulls -o skulls.aar github.com/rootVIII/skulls/skullsebitenbind

// Import the new .aar as a module
// in app/src/main/java/&lt;your username&gt;/MainActivity.java place the following:

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

// add a separate error handling class in app/src/main/java/&lt;your username&gt;/EbitenViewWithErrorHandling.java
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


// Add the below into app/src/main/res/AndroidManifest.xml:
&lt;?xml version="1.0" encoding="utf-8"?&gt;
&lt;manifest xmlns:android="http://schemas.android.com/apk/res/android"
    package="com.&lt;your-username&gt;.skullsmobile"&gt;

    &lt;application
        android:allowBackup="true"
        android:icon="@mipmap/ic_launcher"
        android:label="@string/app_name"
        android:roundIcon="@mipmap/ic_launcher_round"
        android:supportsRtl="true"
        android:theme="@style/Theme.SkullsMobile"&gt;
        &lt;activity android:name=".MainActivity"&gt;
            &lt;intent-filter&gt;
                &lt;action android:name="android.intent.action.MAIN" /&gt;

                &lt;category android:name="android.intent.category.LAUNCHER" /&gt;
            &lt;/intent-filter&gt;
        &lt;/activity&gt;
    &lt;/application&gt;
&lt;/manifest&gt;

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
