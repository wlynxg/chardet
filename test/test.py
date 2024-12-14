import os
import json
import chardet


def detect_and_save_encodings():
    import sys
    import chardet as chardet_lib
    metadata = {
        "python_version": f"{sys.version_info.major}.{sys.version_info.minor}.{sys.version_info.micro}",
        "chardet_version": chardet_lib.__version__,
    }

    base_path = "testdata"
    results = {}

    for root, _, files in os.walk(base_path):
        for file in files:
            file_path = os.path.join(root, file)
            with open(file_path, 'rb') as f:
                rawdata = f.read()
            detection = chardet.detect(rawdata)
            results[file_path] = detection

    output_file = "encoding_results.json"
    with open(output_file, 'w', encoding='utf-8') as json_file:
        json.dump({"metadata": metadata, "results": results}, json_file, indent=4, ensure_ascii=False)


if __name__ == "__main__":
    detect_and_save_encodings()
