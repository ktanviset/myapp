var map;
var golistmarker = [];
var infowindow;
function initMap() {
    map = new google.maps.Map(document.getElementById('map'), {
        center: {lat: 13.7244416, lng: 100.3529108},
        zoom: 10
    });

    infowindow = new google.maps.InfoWindow({
        content: ""
    });
}
function updateMap(){
    $.each(golistmarker, function( index, marker ) {
        marker.setMap(null);
    });
    golistmarker.splice(0, golistmarker.length);
    $('#lolist table tbody').empty();

    let searchString = $("#searchtext").val();//document.getElementById("searchtext").value;
    let url = "/api/GetMakers?keyword=" + searchString;

    let valcountry = $("#countrycode").val();
    if (valcountry != "-1"){
        url += "&countrycode=" + valcountry;
    }
    let valfunction = $("#functioncode").val();
    if (valfunction != "-1"){
        url += "&function=" + valfunction;
    }

    $.get(url, function(data, status){
        console.log(data);

        let bounds = new google.maps.LatLngBounds();

        $.each(data.Makers, function( index, mkdata ) {
            let myLatlng = new google.maps.LatLng(mkdata.latitude,mkdata.longitude);

            if (!(mkdata.latitude == 0 && mkdata.longitude == 0)){
                bounds.extend(myLatlng);
            }

            let myName = mkdata.nameTh;
            if (mkdata.nameEn != null){
                if (myName != ""){
                    myName += "<br>";
                }
                myName += mkdata.nameEn;
            }

            var marker = new google.maps.Marker({
                position: myLatlng,
                title:myName,
            });

            if (!(mkdata.latitude == 0 && mkdata.longitude == 0)){
                marker.setMap(map);
            }

            marker.locode = mkdata.loCode;

            marker.addListener('click', function() {
                let i = index;
                let text = "<div class=\"infowindoes\"><p class=\"text-justify\">" + golistmarker[i].title + "<br>Location Code:" + golistmarker[i].locode + "</p></div>";
                infowindow.setContent(text);
                infowindow.open(map, golistmarker[i]);
            });

            let func0 = "";
            let func1 = "";

            if (mkdata.Func1 != ""){
                if (mkdata.Func1 == "0"){
                    func0 = "Y";
                } else {
                    func0 = "N";
                }

                if (mkdata.Func1 == "1"){
                    func1 = "Y"
                } else {
                    func1 = "N"
                }
            }

            let rowtable = 
            "<tr>"+
                "<th scope=\"row\">"+(index+1)+"</th>"+
                "<td>"+myName+"</td>"+
                "<td>"+mkdata.loCode+"</td>"+
                "<td>"+mkdata.FullCountry+"</td>"+
                "<td>"+func0+"</td>"+//func0
                "<td>"+func1+"</td>"+//func1
                "<td>"+mkdata.Func2+"</td>"+
                "<td>"+mkdata.Func3+"</td>"+
                "<td>"+mkdata.Func4+"</td>"+
                "<td>"+mkdata.Func5+"</td>"+
                "<td>"+mkdata.Func6+"</td>"+
                "<td>"+mkdata.Func7+"</td>"+
                "<td>"+mkdata.Func8+"</td>"+
                "<td>"+mkdata.TruckAmount+"</td>"+
            "</tr>";

            $('#lolist table tbody').append(rowtable);
    
            golistmarker.push(marker);
        });

        map.fitBounds(bounds);
        map.panToBounds(bounds);   
    });
}
$(function() {
    let urlmc = "/api/GetMasterCountry";
    $.get(urlmc, function(data, status){
        console.log(data);

        $.each(data.Masters, function( index, master ) {
            let rowtable = "<option value=\"" + master.val + "\">" + master.display + "</option>";
            $('#countrycode').append(rowtable);
        });
    });

    let urlmf = "/api/GetMasterFunction";
    $.get(urlmf, function(data, status){
        console.log(data);

        $.each(data.Masters, function( index, master ) {
            let rowtable = "<option value=\"" + master.val + "\">" + master.val + " " + master.display + "</option>";
            $('#functioncode').append(rowtable);
        });
    });
});