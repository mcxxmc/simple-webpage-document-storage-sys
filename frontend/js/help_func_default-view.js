const urlDefaultView = "http://localhost:8080/default-view"
const dirIdPrefix = "dir_"
const fileIdPrefix = "file_"
const tab ="&nbsp&nbsp&nbsp&nbsp"
const defaultUserId = "0"

/**
 * returns an id for directory
 * @param {string} dirId 
 * @returns {string}
 */
function makeDirId(dirId) {
    return dirIdPrefix + dirId
}

/**
 * returns an id for file
 * @param {string} fileId 
 * @returns {string}
 */
function makeFileId(fileId) {
    return fileIdPrefix + fileId
}

/**
 * returns the indention depending on the level
 * @param {number} level the directory level or the file level
 * @returns {string}
 */
function levelIndention(level) {
    if (level == 0) {
        return ""
    } else if (level == 1) {
        return tab;
    } else {
        return tab + levelIndention(level - 1)
    }
}

/**
 * append the directory
 * @param {string} dirId 
 * @param {string} dirName 
 * @param {number} dirLevel 
 */
function appendDir(dirId, dirName, dirLevel) {
    $("#div_directories").append("<p 'id'=" + makeDirId(dirId) + "class='directories'>" + levelIndention(dirLevel) + dirName + "</p><br>");
}

/**
 * append the file
 * @param {string} fileId 
 * @param {string} fileName 
 * @param {number} fileLevel 
 */
function appendFile(fileId, fileName, fileLevel) {
    $("#div_directories").append("<a 'id'=" + makeFileId(fileId) + 
    " href='javascript:void(0)' class='files' onClick='clickFile(" + fileId + ")'>" + levelIndention(fileLevel) + fileName + "</a><br>");
}

/**
 * append the whole structure
 * @param {JSON} dirs 
 * @param {string} cur the current directory id to append
 * @param {number} dep the current depth of recursion
 * @param {number} maxDep the max depth of recursion
 */
function recursivelyAppend(dirs, cur, dep, maxDep=24) {
    if (dep >= maxDep) {
        return
    }
    try {
        let dir = dirs[cur];
        if (dir["dir"] == true) {
            let children = dir["children"];
            appendDir(dir["id"], dir["name"], dir["level"]);  // append the current dir
            if (children.length > 0) {  // append its children
                for (let i = 0; i < children.length; i ++) {
                    recursivelyAppend(dirs, children[i], dep + 1);
                }
            }
        } else {
            appendFile(dir["id"], dir["name"], dir["level"]);
        } 
    } catch(error) {
        console.error(error);
    }
}

/**
 * display the JSON data (which contains directory info) on screen 
 * @param {JSON} data 
 *     "tops"
 *     "dirs"
 */
function dispalyDirectories(data) {
    let top = data["top"];
    let dirs = data["dirs"];
    recursivelyAppend(dirs, top, 0);
}

/**
 * display the JSON data (which contains file info) on screen
 * @param {JSON} data 
 *     "file_name"
 *     "content"
 */
function displayFile(data) {
    // clean old data
    $("#div_txt").empty();
    
    // display new data
    $("#div_txt").append("<h>" + data["file_name"] + "</h><br>");
    $("#div_txt").append(data["content"]);
}

/**
 * asks the backend for the file object
 * @param {number} fid the file id
 */
function clickFile(fid) {
    $.ajax({
        type: "POST",
        url: urlDefaultView,
        data: JSON.stringify({"user": defaultUserId, "fid": fid.toString()}),
        success: function(data) {
            displayFile(data);
        },
        error: function(data) {
            alert("error");
        }
    })
}

$(document).ready(function(){

    $.ajax({
        type: "GET",
        url: urlDefaultView,
        success: function(data) {
            dispalyDirectories(data);
        },
        error: function(data) {
            alert("error");
        }
    });
})