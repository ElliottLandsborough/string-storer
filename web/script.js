function getPosts() {
    var xhr = new XMLHttpRequest();

    // Setup our listener to process compeleted requests
    xhr.onreadystatechange = function() {

        // Only run if the request is complete
        if (xhr.readyState !== 4) return;

        // Process our return data
        if (xhr.status >= 200 && xhr.status < 300) {
            result = JSON.parse(xhr.responseText);

            for (var i = 0; i < result.length; i++) {
                setTitle(result[i].Title);
                // Api only returns one post at the moment
                break;
            }
        } else {
            // What to do when the request has failed
            console.log('error', xhr);
        }
    };

    xhr.open("GET", "/api/posts", true);
    xhr.send();
}

getPosts();

var formElement = document.getElementById("form");

formElement.addEventListener("submit", function(evt) {
    evt.preventDefault();
    submitForm();
});

function submitForm() {
    var titleInput = document.getElementById("title-input");
    var title = titleInput.value;

    // limit input to 100 chars
    if (title.length > 100) {
        title = title.slice(0, 100)
    }

    postTitleToApi(title);
}

function postTitleToApi(title) {
    var xhr = new XMLHttpRequest();

    xhr.onreadystatechange = function() {

        // Only run if the request is complete
        if (xhr.readyState !== 4) return;

        // Process our return data
        if (xhr.status >= 200 && xhr.status < 300) {
            result = JSON.parse(xhr.responseText);
            setTitle(result.Title);
        } else {
            // What to do when the request has failed
            console.log('error', xhr);
        }
    };

    xhr.open("POST", "/api/posts");
    xhr.setRequestHeader("Content-Type", "application/json;charset=UTF-8");
    xhr.send(JSON.stringify({
        Title: title
    }));
}

function setTitle(title) {
    document.getElementById("title").innerHTML = 'The saved string is ' + title;
}