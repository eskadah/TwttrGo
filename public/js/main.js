var SearchResult = Backbone.Model.extend({
    validate: function (attrs) {
        if (!attrs.title) return "There must be a title"
    }
});

var SearchResults = Backbone.Collection.extend({
    initialize: function(models, options){
        this.query = options.query
    },
    model: SearchResult,
    url: function() {
        return "/search?query=" + this.query
    }
});

var SearchResultsList = Backbone.View.extend({
    el: "ul#results",
    model: SearchResults,
    render: function () {
        var self = this;
        self.model.each(function (searchResult) {
            var resultView = new SearchResultView({model: searchResult});
            self.$el.append(resultView.render().$el)
        });
        return self
    }
});

var SearchResultView = Backbone.View.extend({
    el: function(){ return $("<li>", {class:"list-group-item" }) },
    events: {
        "click button":function(){
            this.$el.remove()
        }
    },
    render: function () {
        var closeButton = '<button type="button" class="close">' + "<span >&times;</span>" +"</button>";
        this.$el.html(this.model.get("title") + closeButton);
        return this
    }
});

var RestrictedAlert = Backbone.View.extend({
        template: "<div class='alert alert-warning'> <button type='button' class='close' >×</button> Please enter <strong>consonants</strong> only! </div>",
        render: function () {
            this.$el.html(this.template);
            return this
        },
        events: {
            "click button":function() {
                this.$el.empty()
            }
        }
    });

var FormView = Backbone.View.extend({
    events: {
        "submit": "onSubmit",
        "keyup .form-control": "onKeyUp"
    },

    onKeyUp: function(){
        if(this.$el.find("input[type=text]").val().length === 0) {
            $("ul#results").empty()
        }
    },

    onSubmit: function (e) {
        e.preventDefault();
        $("#results").empty();
        var query = this.$el.find("input[type=text]").val();

        if (this._hasRestricted(query)) {
            new RestrictedAlert({ el: "#error" }).render()
        } else {
             new SearchResults(null, {query: query}).fetch({
                success: function (model, response) {
                    //model is search results collection
                    //response is actual results data
                    model.reset();
                    if(response) {
                        response.forEach(function (result) {
                            model.add(new SearchResult({title: result}))
                        });
                    }else{
                        model.add(new SearchResult({title: "No Results"}))
                    }

                    var resultsList = new SearchResultsList({ model: model});
                    resultsList.render();
                },

                error: function (model, response) {
                    console.log("errored");
                    console.log(response)
                }
            })
        }
    },

    _hasRestricted: function (query) {
        //TODO: Use a regex here
        var restricted = [
            "a", "e", "i", "o", "u","0","1","2","3","4","5","6","7","8","9"
        ];
        return _.intersection(restricted, query.split("")).length !== 0
    }
});

$(document).ready(function () {
   var myForm = new FormView({el: "#search-form"})
});
