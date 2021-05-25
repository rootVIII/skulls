### Skulls is simple Columns-like strategy game developed in Golang with the Ebiten library (for Android)


All development/debugging was done with the <b>gomobile</b> tool and <b>adb</b>.


<img src="https://images2.imgbox.com/a6/ab/4hlQKK3q_o.png" alt="ex2"/>


###### Run in a local testing environment:

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


<img src="https://images2.imgbox.com/1d/4c/i9yuv83m_o.png" alt="ex1"/></a>


This was developed on macOS Big Sur.
<hr>
<b>Author: rootVIII  2021</b>
<br><br>
