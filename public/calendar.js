	$(function() {

      var start = moment().startOf('month');
      var end = moment();

      function cb(start, end) {
          $('#startperiod').val(start.format('DD.MM.YYYY'));
          $('#endperiod').val(end.format('DD.MM.YYYY'));
      }

      $('#reportrange').daterangepicker({
					autoApply:true,
          startDate: start,
          endDate: end,
          showDropdowns: true,
          alwaysShowCalendars:true,
          ranges: {
             'Today': [moment(), moment()],
             'Yesterday': [moment().subtract(1, 'days'), moment().subtract(1, 'days')],
             'Last 7 Days': [moment().subtract(6, 'days'), moment()],
             'Last 30 Days': [moment().subtract(29, 'days'), moment()],
             'This Month': [moment().startOf('month'), moment().endOf('month')],
             'Last Month': [moment().subtract(1, 'month').startOf('month'), moment().subtract(1, 'month').endOf('month')]
          }
      }, cb);

      cb(start, end);

  });

