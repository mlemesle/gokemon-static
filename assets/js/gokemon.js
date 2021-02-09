$( document ).ready(function() {
    $("#generationSelect").on('change', function(e){
        var showedDescription = $('#descriptionsContainer .collapse\\.show');
        showedDescription.removeClass('collapse.show');
        showedDescription.addClass('collapse');
        var targetDescription = $($(this).val());
        targetDescription.removeClass('collapse');
        targetDescription.addClass('collapse.show');
    });
});