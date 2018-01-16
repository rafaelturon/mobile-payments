$('document').ready(function() {
    var getBalance = $('#btn-balance');
    getBalance.click(function() {
        callAPI('/api/balance');
      });    
});

function callAPI(endpoint) {
    var accessToken = localStorage.getItem('access_token');
  
    var headers;
    if (accessToken) {
      headers = { Authorization: 'Bearer ' + accessToken };
    }
  
    $.ajax({
      url: endpoint,
      headers: headers
    })
      .done(function(result) {
        $('#balance-view h2').text("Balance: " + result.amount);
      })
      .fail(function(err) {
        $('#balance-view h2').text('Request failed: ' + err.statusText);
      });
  }