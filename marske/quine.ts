import * as fs from "fs";

const pythonQuine = `
s = 's = {!r}\\nprint(s.format(s))\\n'
print(s.format(s))
`;

const rustQuine = `
fn main() {
    let s = "fn main() {\\n    let s = {!r};\\n    println!(s, s);\\n}\\n";
    println!("{}, {}", s, s);
}
`;

const writeToFile = (filename: string, content: string) => {
  fs.writeFileSync(filename, content);
};

const main = () => {
  const typescriptQuine = `
const pythonQuine = \`
${pythonQuine.replace(/`/g, "\\`")}
\`;

const rustQuine = \`
${rustQuine.replace(/`/g, "\\`")}
\`;
`;

  writeToFile("typescript_quine.ts", typescriptQuine);
  writeToFile("python_quine.py", pythonQuine);
  writeToFile("rust_quine.rs", rustQuine);
};

main();
