import zipfile, argparse, os, subprocess, shutil

TEST_DIR = './extracted_test_files/'
TEST_OUT = './test_process.out'

if __name__ == '__main__':
    parser = argparse.ArgumentParser()
    parser.add_argument('-f', '--file', type=str, required=True, help='Golang programm to test')
    parser.add_argument('-t', '--test', type=str, required=True, help='.zip archive with tests and answers')
    parser.add_argument('-n', '--nonstop', action='store_true', default=False)
    parser.add_argument('-s', '--silent', action='store_true', default=False)    
    args = parser.parse_args()

    with zipfile.ZipFile(args.test, 'r') as zip_ref:

        if args.file.endswith('.go'):
            args.file = args.file[:-3]

        if os.path.exists('./' + args.file+'.go'):
            comleted_proc=subprocess.run(['go', 'build', '-o', args.file, './' + args.file + '.go']) 
            if comleted_proc.returncode != 0:
                print("Build error")
                exit()
        else:
            print("No such file or directory: " + args.file + ".go")
                

        if os.path.exists(TEST_DIR):
            shutil.rmtree(TEST_DIR)
        zip_ref.extractall(TEST_DIR)

    for root, dirs, files in os.walk(TEST_DIR):
        for file in sorted([f for f in files if not f.endswith('.a')], key=lambda x: int(x)):
            with open(TEST_DIR + file, 'r') as f:
                with open(TEST_OUT, 'w') as out:
                    subprocess.run(['./' + args.file], stdin=f, stdout=out)
                    with open(TEST_OUT, 'r') as out:
                        with open(TEST_DIR + file + '.a', 'r') as a:
                            result = out.read()
                            expected = a.read()
                            if result == expected:
                                print(f"Тест {file} прошел успешно")
                            else:
                                print(f"Тест {file} провален")
                                if not args.silent:
                                    result_lines = result.splitlines()
                                    expected_lines = expected.splitlines()
                                    for line_number in range(min(len(result_lines), len(expected_lines))):
                                        if result_lines[line_number] != expected_lines[line_number]:
                                            print(f"Expected:\n{expected_lines[line_number]}\nbut got:\n{result_lines[line_number]}")
                                            if not args.nonstop:
                                                exit()
                                if not args.nonstop:
                                    exit()

    if os.path.exists(TEST_DIR):
        os.remove(TEST_OUT)
        shutil.rmtree(TEST_DIR)
