# photo_validator
Helps to identify corrupt images (raw, jpg)

## How to Use

The `photo_validator` tool validates image files (JPEG or RAW `.raf`) and can optionally restore invalid files from a backup source. Below are detailed instructions for usage, based on the available flags and the application's logic.

---

### 1. Basic Usage

To run the application, use the following command:

```sh
photo_validator -f <image_file_or_directory>
```

- The `-f` flag is **required** and specifies the path to a single image file or a directory containing images.

---

### 2. Flags

| Flag         | Type    | Required | Description                                                                                  |
|--------------|---------|----------|----------------------------------------------------------------------------------------------|
| `-f`         | string  | Yes      | Path to the image file or directory to validate.                                             |
| `-s`         | string  | No       | Source directory (backup location) to download and replace invalid images.                   |
| `-p`         | string  | No       | Prefix to ignore in the source path when constructing the backup file path.                  |
| `-a`         | bool    | No       | Ask for confirmation before replacing invalid files (default: `true`).                       |

---

### 3. Examples

#### Validate a Single Image

```sh
photo_validator -f /path/to/image.jpg
```

#### Validate All Images in a Directory

```sh
photo_validator -f /path/to/images/
```

#### Validate and Restore from Backup

If you want the application to attempt restoring invalid images from a backup location:

```sh
photo_validator -f /path/to/images/ -s /mnt/backup/
```

- If an image is invalid, the tool will try to download the replacement from `/mnt/backup/`.

#### Ignore a Prefix in the Source Path

If your backup directory structure differs and you want to ignore a prefix in the file path:

```sh
photo_validator -f /photos/2025/ -s /mnt/backup/ -p /photos/
```

- This will strip `/photos/` from the path when constructing the backup file path.

#### Disable Confirmation Prompt

By default, the tool asks for confirmation before replacing files. To disable this prompt:

```sh
photo_validator -f /path/to/images/ -s /mnt/backup/ -a=false
```

---

### 4. Output

- The tool logs the number of files checked and how many were re-created from backup.
- If a file is invalid and no backup source is provided, the tool will exit with an error.

---

### 5. Notes

- Only `.jpg`, `.jpeg`, and `.raf` files are supported.
- If a directory is provided, all files within will be validated.
- The tool requires read/write permissions for the target files and directories.

---

**Example Full Command:**

```sh
photo_validator -f /photos/2025/ -s /mnt/backup/ -p /photos/ -a=false
```

This will:
- Validate all images in `/photos/2025/`
- Attempt to restore any invalid images from `/mnt/backup/`
- Ignore the `/photos/` prefix when looking up backup files
- Replace files without asking for confirmation

---

**Tip:**  
Run `photo_validator -h` to see all available options and flags.
