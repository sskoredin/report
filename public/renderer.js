// This file is required by the index.html file and will
// be executed in the renderer process for that window.
// No Node.js APIs are available in this process because
// `nodeIntegration` is turned off. Use `preload.js` to
// selectively enable features needed in the rendering
// process.
$("#form").submit(function(e) {
		var start=$('#startperiod').val();
		var end=$('#endperiod').val();

		var url="http://172.28.63.29:8350/report?start="+start+'&end='+end
			var jqxhr = $.get( url, function() {
      })
        .done(function() {
          alert( "success" );
        })
        .fail(function() {
          alert( "error" );
        })
      e.preventDefault();
  });