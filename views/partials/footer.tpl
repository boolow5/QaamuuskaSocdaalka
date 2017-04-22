<footer class="dark-blue-color">
  <div class="container footer-content min-height">
    <div class="row">

      <div class="col-sm-2">
        <h1 class="column-header"><a href="#">{{i18n $.Lang "categories"}}</a></h1>
        <br>
        <ul class="">
          {{range $index, $val := .Categories}}
          <li class="">
            <span class="badge">{{$val.PostsCount}}</span>
            {{i18n $.Lang $val.Name}}
          </li>
          {{end}}
        </ul>
      </div>
      <div class="col-sm-5">
        <h1 class="column-header"><a href="#">{{i18n $.Lang "about"}}</a></h1>
        <br>
        <p>
          Qaamuuska waa bog hadafkiisu yahay inuu dadweynaha ka haqabtiro
          warbixinada la xiriira safarka, safaaradaha, garoomada, baasaboorada,
          xuduudaha, dadweynaha, dhaqamada, sharciyada, wadamada laga necebyahay
          Soomaalida, kuwa laga jecelyahay iyo Goobaha dalxiiska ku haboon.
        </p>
      </div>
      <div class="col-sm-2">

      </div>

      <div class="col-sm-3">
        <h1 class="column-header"><a href="#">{{i18n $.Lang "contact us"}}</a></h1>
        <br>
        <ul>
          <li>Facebook: <a href="#">@Qaamuuska</a></li>
          <li>Twitter: <a href="#">@Qaamuuska</a></li>
          <li>Youtube: <a href="#">@Qaamuuska</a></li>
          <li>Instagram: <a href="#">@Qaamuuska</a></li>
        </ul>
      </div>

    </div>
  </div>

  <div class="copyright">
    <div class="center-text">All Rights Reserved &copy; 2017 Mahdi Bolow</div>
  </div>
</footer>
