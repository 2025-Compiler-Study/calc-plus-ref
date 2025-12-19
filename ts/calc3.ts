import * as fs from 'fs';
import { CharStreams, CommonTokenStream } from 'antlr4ts';
import { CalcPlusLexer } from './CalcPlusLexer';
import { CalcPlusParser } from './CalcPlusParser';
import { Calculator } from './Calculator';

function readSync(): string {
    const BUF_SIZE = 1024;
    const buf = Buffer.alloc(BUF_SIZE);
    let totalStr = '';

    try {
        const bytesRead = fs.readSync(0, buf, 0, BUF_SIZE, null);
        if (bytesRead === 0) return '';
        totalStr = buf.toString('utf8', 0, bytesRead).trim();
    } catch (e) {
        return '';
    }
    return totalStr;
}

function main() {
    const args = process.argv.slice(2);
    if (args.length < 1) {
        console.log("Usage: ts-node calc3.ts <file_path>");
        return;
    }

    const inputCode = fs.readFileSync(args[0], 'utf-8');

    const inputStream = CharStreams.fromString(inputCode);

    const lexer = new CalcPlusLexer(inputStream);
    const tokenStream = new CommonTokenStream(lexer);
    const parser = new CalcPlusParser(tokenStream);

    const tree = parser.calc3();

    const visitor = new Calculator(
        () => readSync(),
        (msg) => console.log(msg)
    );

    visitor.visit(tree);
}

main();