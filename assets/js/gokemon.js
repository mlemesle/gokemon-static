$( document ).ready(function() {
    $("#generationSelect").on('change', function(e){
        var showedDescription = $('#descriptionsContainer .collapse\\.show');
        showedDescription.removeClass('collapse.show');
        showedDescription.addClass('collapse');
        var targetDescription = $($(this).val());
        targetDescription.removeClass('collapse');
        targetDescription.addClass('collapse.show');
    });

    var firstDescription = $("#descriptionsContainer .collapse").first();
    firstDescription.removeClass('collapse');
    firstDescription.addClass('collapse.show');
});