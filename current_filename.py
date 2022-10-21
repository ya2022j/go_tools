def __current_filename():
    from inspect import currentframe, getframeinfo
    frameinfo = getframeinfo(currentframe())
    current_filename = os.path.basename(frameinfo.filename)
    return current_filename
