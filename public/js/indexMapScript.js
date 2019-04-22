var map;
var markers = [];
var infowindow;
function initMap() {
    map = new google.maps.Map(document.getElementById('map'), {
        center: {lat: 13.7244416, lng: 100.3529108},
        zoom: 10
    });

    infowindow = new google.maps.InfoWindow({
        content: ""
    });

    $.get("/api/GetMakers", function(data, status){
        console.log(data);

        $.each(data.Makers, function( index, mkdata ) {
            let myLatlng = new google.maps.LatLng(mkdata.latitude,mkdata.longitude);
            let myName = mkdata.nameTh;
            if (mkdata.nameEn != null){
                myName += "<br>" + mkdata.nameEn;
            }

            var marker = new google.maps.Marker({
                position: myLatlng,
                title:myName,
                map: map,
            });

            marker.addListener('click', function() {
                let i = index;
                let text = "<div class=\"infowindoes\"><p class=\"text-justify\">"+markers[i].title+"</p></div>";
                infowindow.setContent(text);
                infowindow.open(map, markers[i]);
            });
    
            markers.push(marker);
        });
    });
}