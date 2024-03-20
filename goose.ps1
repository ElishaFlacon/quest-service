#! /bin/pwsh

Get-Content .\.env | ForEach-Object {
    $name, $value = $_.split('=')
    if (!([string]::IsNullOrWhiteSpace($name))) {
        Write-Output "$name=$value"
        Set-Content env:\$name $value
    }

}

goose -dir migrations $c
