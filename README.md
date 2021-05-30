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
ebitenmobile bind -target android -javapkg com.&lt;your username&gt;.skulls -o skulls.aar github.com/rootVIII/skulls/skullsebitenbind
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
