' Variable and Constant Declarations
Const FOF_CREATEPROGRESSDLG = &H0&
Const OverwriteExisting = True
Const ForWriting = 2
dim strDate, strBackupSource, strBackupDest
' Change Backup Source and Destination
strBackupSource = "C:\scripts\test1\*.*"
strBackupDest = "C:\scripts\test2"
GetDate

' Object Assignments and File/Folder Creation

Set objFSO = CreateObject("Scripting.FileSystemObject")
Set objFile = objFSO.CreateTextFile("C:\abc\backup.log", ForWriting, True)
Set objShell = CreateObject("Shell.Application")
strBackupDest = strBackupDest & "\" & strdate
objFSO.CreateFolder(strBackupDest)
Set objFolder = objShell.NameSpace(strBackupDest)
Set BackupArchive = objFSO.GetFolder(strBackupDest)
Set collFolders = BackupArchive.SubFolders

' Code Body
For Each subFol in collFolders
	ymd = subFol.DateCreated
	If DateDiff("d", ymd, now) > 30 Then objFSO.DeleteFolder subFol, True End IF
Next
objFile.Write "Backup Started " & now
objFolder.CopyHere strBackupSource, FOF_CREATEPROGRESSDLG
objFile.Write vbCrLf & "Backup Ended " & now
objFile.Close
WScript.Quit

' SubRoutines
Sub GetDate
	If Len(Day(Now)) = 1 Then
		strDay = "0" & Day(Now)
	Else
		strDay = Day(Now)
	End If
	If Len(Month(Now)) = 1 Then
		strMonth = "0" & Month(Now)
	Else
		strMonth = Month(Now)
	End If
	strYear = Year(Now)
	If Len(Hour(Now)) = 1 Then
		strHour = "0" & Hour(Now)
	Else
		strHour = Hour(Now)
	End If
	If Len(Minute(Now)) = 1 Then
		strMin = "0" & Minute(Now)
	Else
		strMin = Minute(Now)
	End If
	strDate = strYear & strMonth & strDay & strHour & strMin
End Sub