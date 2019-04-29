var markersData = [];

$( document ).ready(function() {
    console.log( "ready!" );
});

function clearForm(){

}

function updateForm(i){
    //alert("test" + i);
    $("#lo_id").val(markersData[i].id);
    $("#lo_nameth").val(markersData[i].nameTh);
    $("#lo_nameen").val(markersData[i].nameEn);
    $("#lo_latitude").val(markersData[i].latitude);
    $("#lo_longitude").val(markersData[i].longitude);
    $("#lo_locode").val(markersData[i].loCode);
}

function search(){
    markersData.splice(0, markersData.length);
    $('#lolist table tbody').empty();

    let searchString = $("#searchtext").val();
    let url = "/api/GetMakers?keyword=" + searchString;
    $.get(url, function(data, status){
        console.log(data);

        $.each(data.Makers, function( index, mkdata ) {
            mkdata.index = index;
            markersData.push(mkdata);

            let myName = mkdata.nameTh;
            if (mkdata.nameEn != null){
                myName += "<br>" + mkdata.nameEn;
            }

            let rowtable = "<tr onclick=\"updateForm("+index+")\">"+
                "<th scope=\"row\">"+(index+1)+"</th>"+
                "<td>"+myName+"</td>"+
            "</tr>";

            $('#lolist table tbody').append(rowtable);
        });
    });
}