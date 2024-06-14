
# Loop over all files in examples and generated into /images
Get-ChildItem -Path examples -Filter *.* | ForEach-Object {
    $output = "images/$($_.BaseName).png"
    # if image already exists, skip
    if (Test-Path $output) {
        return
    }
    Write-Host "Generating '$output'"

    freeze -c user $_.FullName -o $output
}